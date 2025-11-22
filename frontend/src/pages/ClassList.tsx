import { useEffect, useState } from "react";
import { deleteClass, getAllClass } from "../server/class";
import type { Class } from "../types/type";
import { useNavigate } from "react-router-dom";

export default function ClassList() {
  const [classes, setClasses] = useState<Class[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const [showModal, setShowModal] = useState(false);
  const [newClass, setNewClass] = useState({ name: "", category: "" });

  const navigate = useNavigate();

  const handleCreate = () => {
    if (!newClass.name || !newClass.category) return;
    setClasses([...classes, { id: Date.now(), ...newClass }]);
    setNewClass({ name: "", category: "" });
    setShowModal(false);
  };

  const handleDelete = async (id: number) => {
    try {
      await deleteClass(id);
      alert("success delete class");
    } catch (err) {
      alert("Failed to delete classes");
      console.log(err);
    } finally {
      window.location.reload(); // reload seluruh halaman
    }
  };

  useEffect(() => {
    const fetchData = async () => {
      try {
        const data = await getAllClass();
        console.log(data);
        setClasses(data);
      } catch (err) {
        setError("Failed to fetch classes");
        console.log(err);
      } finally {
        setLoading(false);
      }
    };

    fetchData();
  }, []);

  if (loading) return <div className="text-black">Loading...</div>;

  if (error)
    return (
      <div className="flex justify-center items-center h-screen">
        <p className="text-2xl font-bold text-red-500">{error}</p>
      </div>
    );

  return (
    <div>
      <div className="flex justify-between items-center mb-6">
        <h1 className="text-3xl font-bold text-gray-800">Classes</h1>
        <button
          className="btn btn-primary"
          onClick={() => {
            navigate(`/classes/add`);
          }}
        >
          + Add Class
        </button>
      </div>

      <div className="overflow-x-auto shadow-md rounded-lg bg-white">
        <table className="table-auto w-full text-left border-collapse">
          <thead className="bg-gray-100 text-black">
            <tr>
              <th className="px-6 py-3 border-b">No</th>
              <th className="px-6 py-3 border-b">Name</th>
              <th className="px-6 py-3 border-b">Category</th>
              <th className="px-6 py-3 border-b w-36">Actions</th>
            </tr>
          </thead>
          <tbody>
            {Array.isArray(classes) &&
              classes.map((c, index) => (
                <tr key={c.id} className="hover:bg-gray-50">
                  <td className="px-6 py-4 text-black border-b">{index + 1}</td>
                  <td className="px-6 py-4 text-black border-b">{c.name}</td>
                  <td className="px-6 py-4 text-black border-b">
                    {c.category}
                  </td>
                  <td className="px-6 py-4 text-black border-b flex gap-2">
                    <button
                      className="btn btn-sm btn-ghost"
                      onClick={() => {
                        navigate(`/classes/` + c.id);
                      }}
                    >
                      <h5 className="text-white">Detail</h5>
                    </button>
                    <button
                      className="btn btn-sm btn-error"
                      onClick={() => handleDelete(c.id)}
                    >
                      <h5 className="text-white">Delete</h5>
                    </button>
                  </td>
                </tr>
              ))}
          </tbody>
        </table>
      </div>
    </div>
  );
}
