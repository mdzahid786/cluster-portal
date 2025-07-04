import { useState } from "react";

export default function ClusterCard({ cluster, isAdmin, onUpdate }) {
  const [editing, setEditing] = useState(false);
  const [servers, setServers] = useState(cluster.servers);
  const [loading, setLoading] = useState(false);

  const handleEdit = () => setEditing(true);
  const handleCancel = () => {
    setEditing(false);
    setServers(cluster.servers);
  };
  const handleSave = async () => {
    setLoading(true);
    await onUpdate(cluster.id, servers);
    setEditing(false);
    setLoading(false);
  };

  return (
    <div className="cluster-card">
      <h3>{cluster.name}</h3>
      <div className="servers-row">
        {editing ? (
          <>
            <input
              type="number"
              min="0"
              value={servers}
              onChange={(e) => setServers(Number(e.target.value))}
              disabled={loading}
            />
            <button onClick={handleSave} disabled={loading}>
              &#10004;
            </button>
            <button onClick={handleCancel} disabled={loading}>
              &#10006;
            </button>
          </>
        ) : (
          <>
            <span className="servers-count">{cluster.servers}</span>
            {isAdmin && (
              <button
                className="edit-btn"
                onClick={handleEdit}
                title="Edit servers"
              >
                &#9998;
              </button>
            )}
          </>
        )}
      </div>
    </div>
  );
}
