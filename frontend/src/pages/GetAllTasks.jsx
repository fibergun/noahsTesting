import {useSession} from "../routers/Session.jsx";
import {useEffect, useState} from "react";

function GetAllTasks(){
    const [tasks, setTasks] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);
    const {session, loading: sessionLoading} = useSession();
    const [taskDetails, setTaskDetails] =useState({})
    const [points, setPoints] = useState({})

    useEffect(() => {
        if (sessionLoading || !session) return;

        async function handleRequest() {
            try {
                const response = await fetch(`/tasks/list?userID=${session.userID}`, {
                    method: 'GET',
                    headers: {'Content-Type': 'application/json'},
                });

                if (!response.ok) {
                    const message = await response.text();
                    setError(message);
                    return;
                }

                const data = await response.json();
                setTasks(data.tasks);
                setPoints(data.points)
            } catch (err) {
                setError(err.message);
            } finally {
                setLoading(false);
            }
        }

        handleRequest();

    }, [session, sessionLoading]);

    async function completeTask(taskID) {
        try {
            const response = await fetch(`/tasks/complete?taskID=${taskID}`, {
                method: 'POST',
            });

            if (!response.ok) {
                alert(await response.text());
                return;
            }

            setTasks((prev) => prev.filter((task) => task.taskID !== taskID));
        } catch (err) {
            alert("Something went wrong: " + err.message);
        }
    }

    async function getTask(taskID) {
        try {
            const response = await fetch(`/tasks/get?taskID=${taskID}`, {
                method: 'GET',
            });

            if (!response.ok) {
                alert(await response.text());
                return;
            }
            const data = await response.json();
            setTaskDetails((prev) => ({ ...prev, [taskID]: data }));
        } catch (err) {
            alert("Something went wrong: " + err.message);
        }
    }

    useEffect(() => {
        tasks.forEach((task) => {
            if (!taskDetails[task.taskID]) getTask(task.taskID);
        });
    }, [tasks]);

    if (loading) return <p>Loading tasks...</p>;
    if (error) return <p>Error: {error}</p>;

    return (
        <ul>
            <p>
                You have: {points} points!
            </p>
            {tasks.map((task) => (
                <li key={task.taskID}>
                    {taskDetails[task.taskID]?.task ?? 'Loading...'}
                    <button onClick={() => completeTask(task.taskID)}>Complete</button>
                </li>
            ))}
        </ul>
    );
}

export default GetAllTasks;