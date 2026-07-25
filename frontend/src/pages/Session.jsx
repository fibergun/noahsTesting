
import {createContext, useContext, useState, useEffect} from "react";

const SessionContext = createContext(null)

export function SessionProvider({ children }){
const [session, setSession] = useState(null);
const [loading, setLoading] = useState(true);

useEffect(() =>{
    const stored = localStorage.getItem('session');
    if (stored){
        setSession (JSON.parse(stored));
    }
    setLoading(false);
}, [])

    const login = (userName) => {
    const sessionData = {user: userName, loggedInAt: Date.now()};
    localStorage.setItem('session', JSON.stringify(sessionData));
    setSession(sessionData);
    };

const logout = () => {
    localStorage.removeItem('session');
    setSession(null);
};

 return (
     <SessionContext.Provider value={{session, login, logout, loading}}>
         {children}
     </SessionContext.Provider>
 );
}

export function useSession(){
    const ctx = useContext(SessionContext);
    if (!ctx) throw new Error('useSession must be used within SessionProvider');
    return ctx;
}
