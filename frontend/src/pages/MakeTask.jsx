import {useSession} from "../routers/Session.jsx";
import {useState} from "react";

function MakeTask(){
    const {session} = useSession()
    const [task, setTask] = useState("task")

    async function handleMakeTask(event){
        event.preventDefault()

        const response = await fetch(`/tasks/make`, {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({userID: session.userID, task})
        })

        if (response.ok){
            const data = await response.json()
            alert("succesfully created task: "+ data.task+ " with ID: "+ data.taskID)
        }else{
            const message = await response.text()
            alert(message)
        }
    }

    return(
        <form onSubmit={handleMakeTask}>
            <input value={task} onChange={(event) => setTask(event.target.value)} placeholder="task"/>
            <button type="submit">Submit task</button>
        </form>
    )
}

export default MakeTask