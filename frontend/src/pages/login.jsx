import {useSession} from "./Session.jsx";
import {useState} from "react";

 function LoginForm(){
    const {login} = useSession()
    const [user, setUser] = useState("user")

    async function handleSubmit(event){
        event.preventDefault()

        const response = await fetch('/user', {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({user}),
        })

        if (response.ok){
            const data = await response.json()
            login(data.user)
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