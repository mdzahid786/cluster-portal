import { useState, useEffect } from "react";
import "./App.css";
import Login from "./Login";
import ClusterCard from "./ClusterCard";

function App() {
  const [user, setUser] = useState(null);
  const [clusters, setClusters] = useState([]);
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);
  const [auth, setAuth] = useState(null);

  // Login handler
  const handleLogin = async (username, password) => {
    setError("");
    setLoading(true);
    const basicAuth = "Basic " + btoa(`${username}:${password}`);
    try {
      const res = await fetch("http://localhost:8083/api/me", {
        headers: { Authorization: basicAuth },
      });
      if (!res.ok) throw new Error("Invalid credentials");
      const user = await res.json();
      setUser(user);
      setAuth(basicAuth);
      setLoading(false);
    } catch (e) {
      setError(e.message);
      setLoading(false);
    }
  };

  // Fetch clusters
  useEffect(() => {
    if (!auth) return;
    setLoading(true);
    fetch("http://localhost:8083/api/clusters", {
      headers: { Authorization: auth },
    })
      .then((res) => res.json())
      .then((data) => {
        setClusters(data);
        setLoading(false);
      });
  }, [auth]);

  // Update cluster servers
  const handleUpdate = async (id, servers) => {
    setLoading(true);
    await fetch(`http://localhost:8083/api/cluster/${id}`, {
      method: "PUT",
      headers: {
        Authorization: auth,
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ servers }),
    });
    // Refresh clusters
    const res = await fetch("http://localhost:8083/api/clusters", {
      headers: { Authorization: auth },
    });
    setClusters(await res.json());
    setLoading(false);
  };

  if (!user) {
    return <Login onLogin={handleLogin} error={error} />;
  }

  return (
    <div className="portal-container">
      <header>
        <div className="h1-heading">
          <h1>Cluster Portal</h1>
        </div>
        <div className="user-info">
          Logged in as <b>{user.Username}</b> ({user.Role}){" "}
          <button
            onClick={() => {
              setUser(null);
              setAuth(null);
            }}
          >
            Logout
          </button>
        </div>
      </header>
      {loading && <div className="loading">Loading...</div>}
      <div className="clusters-grid">
        {clusters.map((cluster) => (
          <ClusterCard
            key={cluster.id}
            cluster={cluster}
            isAdmin={user.Role === "admin"}
            onUpdate={handleUpdate}
          />
        ))}
      </div>
    </div>
  );
}

export default App;
