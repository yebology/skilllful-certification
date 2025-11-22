import { useNavigate, useParams } from "react-router-dom";
import { useEffect, useState } from "react";
import type { ClassDetail } from "../types/type";
import { getClassDetail, updateClasss } from "../server/class";

export default function ClassDetail() {
  const { id } = useParams<{ id: string }>();

  const [classDetail, setClassDetail] = useState<ClassDetail>();
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [saving, setSaving] = useState(false);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const data = await getClassDetail(Number(id));
        setClassDetail(data!);
      } catch (err) {
        setError("Failed to fetch classes");
        console.log(err);
      } finally {
        setLoading(false);
      }
    };

    fetchData();
  }, []);

  const navigate = useNavigate();

  const handleSave = async () => {
    if (!classDetail) return;
    setSaving(true);

    let categoryId: number;

    if (classDetail.category === "Desain Grafis") {
      categoryId = 1;
    } else if (classDetail.category === "Pemrograman Dasar") {
      categoryId = 2;
    } else if (classDetail.category === "Editing Video") {
      categoryId = 3;
    } else if (classDetail.category === "Public Speaking") {
      categoryId = 4;
    } else {
      categoryId = 0; // default kalau tidak match
    }

    try {
      await updateClasss(
        classDetail.id,
        classDetail.name,
        categoryId,
        classDetail.description,
        classDetail.instructor
      );
      alert("class updated successfully!");
    } catch (err) {
      console.error(err);
      alert("Failed to update class");
    } finally {
      setSaving(false);
      navigate(`/`);
    }
  };

  if (loading) return <div className="text-black">Loading...</div>;

  if (error)
    return (
      <div className="flex justify-center items-center h-screen">
        <p className="text-2xl font-bold text-red-500">{error}</p>
      </div>
    );

  return (
    <div className="p-6 max-w-xl mx-auto text-black">
      <h1 className="text-2xl font-bold mb-4">Class Detail</h1>
      <div className="space-y-4">
        <div>
          <label className="block mb-1 font-medium">Name</label>
          <input
            className="w-full border border-gray-500 rounded px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-400"
            value={classDetail!.name}
            onChange={(e) =>
              setClassDetail({ ...classDetail!, name: e.target.value })
            }
          />
        </div>

        <div>
          <label className="block mb-1 font-medium">Category</label>
          <select
            className="w-full border border-gray-500 rounded px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-400"
            value={classDetail!.category}
            onChange={(e) =>
              setClassDetail({ ...classDetail!, category: e.target.value })
            }
          >
            <option value="Desain Grafis">Desain Grafis</option>
            <option value="Pemrograman Dasar">Pemrograman Dasar</option>
            <option value="Editing Video">Editing Video</option>
            <option value="Public Speaking">Public Speaking</option>
          </select>
        </div>

        <div>
          <label className="block mb-1 font-medium">Description</label>
          <textarea
            className="w-full border border-gray-500 rounded px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-400"
            rows={4}
            value={classDetail!.description}
            onChange={(e) =>
              setClassDetail({ ...classDetail!, description: e.target.value })
            }
          />
        </div>
        <div>
          <label className="block mb-1 font-medium">Instructor</label>
          <input
            className="w-full border border-gray-500 rounded px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-400"
            value={classDetail!.instructor}
            onChange={(e) =>
              setClassDetail({ ...classDetail!, instructor: e.target.value })
            }
          />
        </div>
        <button
          onClick={handleSave}
          className="btn text-white btn-primary"
        >
          Save
        </button>
      </div>
    </div>
  );
}
