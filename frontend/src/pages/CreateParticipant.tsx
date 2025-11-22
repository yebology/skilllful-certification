import { useNavigate } from "react-router-dom";
import { useState } from "react";
import type { ParticipantDetail } from "../types/type";
import { createParticipant } from "../server/participant"; // buat function API

export default function CreateParticipant() {
  const [participant, setParticipant] = useState<ParticipantDetail>({
    id: 0,
    name: "",
    email: "",
    gender: "",
    phoneNumber: "",
    birthDate: "",
  });
  const [saving, setSaving] = useState(false);

  const navigate = useNavigate();

  const handleSave = async () => {
    if (!participant) return;
    setSaving(true);

    const genderId = participant.gender == "Laki-Laki" ? 1 : 2;

    try {
      await createParticipant(
        participant.name,
        participant.email,
        genderId,
        participant.birthDate,
        participant.phoneNumber
      );
      alert("Participant created successfully!");
    } catch (err) {
      console.error(err);
      alert("Failed to create participant");
    } finally {
      setSaving(false);
      navigate(`/participants`);
    }
  };

  return (
    <div className="p-6 max-w-xl mx-auto text-black">
      <h1 className="text-2xl font-bold mb-6">Create Participant</h1>
      <div className="space-y-4">
        {/* Name */}
        <div>
          <label className="block mb-1 font-medium">Name</label>
          <input
            className="w-full border border-gray-500 rounded px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-400"
            value={participant!.name}
            onChange={(e) =>
              setParticipant({ ...participant!, name: e.target.value })
            }
          />
        </div>

        {/* Email */}
        <div>
          <label className="block mb-1 font-medium">Email</label>
          <input
            type="email"
            className="w-full border border-gray-500 rounded px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-400"
            value={participant!.email}
            onChange={(e) =>
              setParticipant({ ...participant!, email: e.target.value })
            }
          />
        </div>

        {/* Gender */}
        <div>
          <label className="block mb-1 font-medium">Gender</label>
          <select
            className="w-full border border-gray-500 rounded px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-400"
            value={participant!.gender} // sekarang menyimpan "1" atau "2"
            onChange={(e) =>
              setParticipant({ ...participant!, gender: e.target.value })
            }
          >
            <option value="Laki-Laki">Laki-Laki</option>
            <option value="Wanita">Wanita</option>
          </select>
        </div>

        {/* Phone Number */}
        <div>
          <label className="block mb-1 font-medium">Phone Number</label>
          <input
            className="w-full border border-gray-500 rounded px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-400"
            value={participant!.phoneNumber}
            onChange={(e) =>
              setParticipant({ ...participant!, phoneNumber: e.target.value })
            }
          />
        </div>

        {/* Birth Date */}
        <div>
          <label className="block mb-1 font-medium">Birth Date</label>
          <input
            type="date"
            className="w-full border border-gray-500 rounded px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-400"
            value={participant!.birthDate}
            onChange={(e) =>
              setParticipant({ ...participant!, birthDate: e.target.value })
            }
          />
        </div>

        {/* Save button */}
        <button
          type="button"
          className={`text-white mt-4 ${saving ? "loading" : ""}`}
          onClick={handleSave}
          disabled={saving}
        >
          {saving ? "Saving..." : "Save"}
        </button>
      </div>
    </div>
  );
}
