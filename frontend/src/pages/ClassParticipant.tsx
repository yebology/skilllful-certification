import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import {
  assignParticipant,
  deleteClassParticipant,
  getClassParticipant,
} from "../server/classParticipant";
import type { ClassParticipant, Participant } from "../types/type";
import { getAllParticipant } from "../server/participant";

export default function ClassParticipant() {
  const { id } = useParams<{ id: string }>();
  const [participants, setParticipants] = useState<ClassParticipant[]>([]);
  const [allParticipants, setAllParticipants] = useState<Participant[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const [showModal, setShowModal] = useState(false);
  const [selectedParticipantIds, setSelectedParticipantIds] = useState<
    number[]
  >([]);

  // fetch class participants
  useEffect(() => {
    const fetchParticipants = async () => {
      try {
        const data = await getClassParticipant(Number(id));
        setParticipants(data || []);
      } catch (err) {
        console.error(err);
        setError("Failed to fetch participants");
      } finally {
        setLoading(false);
      }
    };
    fetchParticipants();
  }, [id]);

  // fetch all participants, filter out participants already in class
  useEffect(() => {
    const fetchAllParticipants = async () => {
      try {
        const data = await getAllParticipant();
        const filtered = data.filter(
          (p) => !participants.some((cp) => cp.participantId === p.id)
        );
        setAllParticipants(filtered);
      } catch (err) {
        console.error(err);
        setError("Failed to fetch participants");
      }
    };
    fetchAllParticipants();
  }, [participants]);

  const handleDelete = async (participantId: number) => {
    try {
      await deleteClassParticipant(participantId);
      alert("Participant removed successfully");
      window.location.reload();
    } catch (err) {
      console.error(err);
      alert("Failed to remove participant");
    }
  };

  if (loading)
    return (
      <div className="text-black text-center py-10">
        Loading participants...
      </div>
    );

  if (error)
    return (
      <div className="text-red-500 text-center py-10 font-bold">{error}</div>
    );

  return (
    <div className="p-6 mx-auto text-black">
      <div className="flex justify-between items-center mb-6">
        <h1 className="text-3xl font-bold text-gray-800">Class Participants</h1>
        <button className="btn btn-primary" onClick={() => setShowModal(true)}>
          <h5 className="text-white">Assign</h5>
        </button>
      </div>

      <div className="overflow-x-auto shadow-md rounded-lg bg-white">
        <table className="table-auto w-full text-left border-collapse">
          <thead className="bg-gray-100 text-black">
            <tr>
              <th className="px-6 py-3 border-b">No</th>
              <th className="px-6 py-3 border-b">Name</th>
              <th className="px-6 py-3 border-b">Email</th>
              <th className="px-6 py-3 border-b">Gender</th>
              <th className="px-6 py-3 border-b w-36">Actions</th>
            </tr>
          </thead>
          <tbody>
            {participants.length === 0 && (
              <tr>
                <td colSpan={5} className="text-center py-4">
                  No participants found
                </td>
              </tr>
            )}
            {participants.map((p, index) => (
              <tr key={p.id} className="hover:bg-gray-50">
                <td className="px-6 py-4 border-b">{index + 1}</td>
                <td className="px-6 py-4 border-b">{p.name}</td>
                <td className="px-6 py-4 border-b">{p.email}</td>
                <td className="px-6 py-4 border-b">{p.gender}</td>
                <td className="px-6 py-4 border-b flex gap-2">
                  <button
                    className="btn btn-sm btn-error"
                    onClick={() => handleDelete(p.id)}
                  >
                    <h5 className="text-white">Delete</h5>
                  </button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>

      {/* Assign Modal */}
      {showModal && (
        <div className="modal modal-open">
          <div className="modal-box max-w-lg mx-auto">
            <h3 className="font-bold text-xl mb-4">Assign Participant</h3>
            <div className="mb-4">
              <select
                className="w-full border border-gray-300 rounded px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-400"
                value={selectedParticipantIds[0] || ""}
                onChange={(e) =>
                  setSelectedParticipantIds([Number(e.target.value)])
                }
              >
                <option value="">Select participant</option>
                {allParticipants.map((p) => (
                  <option key={p.id} value={p.id}>
                    {p.name} ({p.email})
                  </option>
                ))}
              </select>
            </div>
            <div className="modal-action justify-end">
              <button
                className="btn text-white"
                onClick={() => setShowModal(false)}
              >
                Cancel
              </button>
              <button
                className="btn text-white btn-primary"
                onClick={async () => {
                  if (!selectedParticipantIds[0]) return;
                  try {
                    await assignParticipant(
                      selectedParticipantIds[0],
                      Number(id)
                    );
                    alert("Participant assigned successfully");
                    setShowModal(false);
                    window.location.reload(); // auto refresh
                  } catch (err) {
                    console.error(err);
                    alert("Failed to assign participant");
                  }
                }}
              >
                Assign
              </button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}
