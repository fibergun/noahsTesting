import {useSession} from "../routers/Session.jsx";
import {useEffect, useState} from "react";

function GetAllTasks(){
    const [tasks, setTasks] = useState([])
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);
    const {session} = useSession()

    useEffect(() => {


    async function handleRequest() {
        try {
            const response = await fetch(`/tasks/list?userID=${session.userID}`, {
                method: 'GET',
                headers: {'Content-Type': 'application/json'},
            })

            if (!response.ok) {
                const message = await response.text()
                alert(message)
            }
            const data = await response.json();
            setTasks(data);
        } catch (err) {
            alert("Something went wrong: " + err.message);
        } finally {
            setLoading(false);
        }
    }

        handleRequest();

    }, [])

    if (loading) return <p>Loading tasks...</p>;
    if (error) return <p>Error: {error}</p>;

    return(
        <ul>
            {tasks.map((task) =>(
                <li key={task.id}>{task.name}</li>
            ))}
        </ul>
    );
}

export default GetAllTasks;