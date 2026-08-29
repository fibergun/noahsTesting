import { useSession } from "../routers/Session.jsx";

function capitalize(str) {
  if (!str) return str; // handle empty/undefined safely
  return str.charAt(0).toUpperCase() + str.slice(1);
}

function Home() {
  const { session } = useSession();

  return (
    <div>
      <h1>Welcome {capitalize(session.username)}</h1>
      <p>This is the home page.</p>
    </div>
  );
}

export default Home;
