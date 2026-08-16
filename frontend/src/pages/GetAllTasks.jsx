import {useSession} from "../routers/Session.jsx";
import {useEffect, useState} from "react";

function GetAllTasks(){
    const [tasks, setTasks] = useState([])
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);
    const {session, loading: sessionLoading} = useSession()

    useEffect(() => {
            if(sessionLoading || !session)return;

    async function handleRequest() {
        try {
            const response = await fetch(`/tasks/list?userID=${session.userID}`, {
                method: 'GET',
                headers: {'Content-Type': 'application/json'},
            })

            if (!response.ok) {
                const message = await response.text()
                setError(message);
                return;
            }
            const data = await response.json();
            setTasks(data.tasks);
        } catch (err) {
            alert("Something went wrong: " + err.message);
        } finally {
            setLoading(false);
        }
    }

        handleRequest();

    }, [session, sessionLoading])

    async function completeTask(taskID) {
        try {
            const response = await fetch(`/tasks/complete?taskID=${taskID}`, {
                method: 'POST',
            });
            if (!response.ok) {
                const message = await response.text();
                alert(message)
                return;
            }
            setTasks((prev) => prev.filter((t) => t.taskID !== taskID));
        } catch (err) {
            alert("Something went wrong: " + err.message);
        }
    }

    if (loading) return <p>Loading tasks...</p>;
    if (error) return <p>Error: {error}</p>;

    return(
        <ul>
            {tasks.map((task) => (
                <li key={task.task}>
                    {task.task}
                    <button onClick={() => completeTask(task.taskID)}>Complete</button>
                </li>
            ))}
        </ul>
    );
}

export default GetAllTasks;