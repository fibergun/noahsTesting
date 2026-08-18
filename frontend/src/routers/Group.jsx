import LoginForm from "../pages/login.jsx";
import { useParams } from "react-router-dom";

function Group() {
  const { group } = useParams();

  return (
    <div>
      <p>Logging into group: {group}</p>
      <LoginForm group={group} />
    </div>
  );
}

export default Group;
