import {useSession} from "./Session.jsx";
import {Navigate} from "react-router-dom";


function ProtectedRoute({children}){
    const {session, loading} = useSession()

    if (loading){
        return <p>Loading...</p>;
    }

    if (!session){
        return <Navigate to="/login" replace />;
    }
}