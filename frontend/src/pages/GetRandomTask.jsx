import {useEffect, useState} from "react";
import {useSession} from "../routers/Session.jsx";


function GetRandomTask(){
    const [task, setTask] = useState(null)
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState();
    const {session} = useSession()

    useEffect(() => {
        const controller = new AbortController();

        async function handleRequest() {
            try {
                const response = await fetch(`/tasks/random?userID=${session.userID}`, {
                    method: 'GET',
                    headers: {'Content-Type': 'application/json'},
                    signal: controller.signal,
                })
                if (!response.ok) {
                    const message = await response.text()
                    setError(message)
                    alert(message)
                    return
                }
                const data = await response.json();
                setTask(data);

            } catch (err) {
                if (err.name === 'AbortError'){

                    return
                }
                setError("Something went wrong: " + err.message);
            } finally {
                if(!controller.signal.aborted){
                    setLoading(false);
                }
            }
        }
        handleRequest();

        return () => controller.abort();
    }, [])

    if (loading) return <p>Loading task...</p>;
    if (error) return <p>Error: {error}</p>;
    if (!task) return <p>No task found.</p>;

    return(
        <div>
            <p>{task.task}</p>
        </div>
    );
}

export default GetRandomTask