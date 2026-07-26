import {useSession} from "../routers/Session.jsx";
import {useState} from "react";
import {useNavigate} from "react-router-dom";

 function LoginForm(){
    const {login} = useSession()
    const [user, setUser] = useState("user")
     const navigate = useNavigate()

    async function handleSubmit(event){
        event.preventDefault()

        const response = await fetch('/user', {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({user}),
        })

        if (response.ok){
            const data = await response.json()
            login(data.username, data.userID)
            navigate(`/`)
        }else{
            alert('Login failed')
        }
    }

    return (
        <form onSubmit={handleSubmit}>
            <input value={user} onChange={(event) => setUser(event.target.value)} placeholder="user" />
            <button type="submit">Log in</button>
        </form>
    )
}

export default LoginForm