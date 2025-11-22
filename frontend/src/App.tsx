import { Routes, Route } from "react-router-dom";
import Navbar from "./components/Navbar";
import ClassList from "./pages/ClassList";
import ClassDetail from "./pages/ClassDetail";
import ParticipantList from "./pages/ParticipantList";
import ParticipantDetail from "./pages/ParticipantDetail";
import CreateClass from "./pages/CreateClass";
import CreateParticipant from "./pages/CreateParticipant";
import ClassParticipant from "./pages/ClassParticipant";
import ParticipantClass from "./pages/ParticipantClass";

function App() {
  return (
    <div className="min-h-screen w-screen bg-gray-100">
      <Navbar />
      <div className="max-w-7xl mx-auto px-6 py-8">
        <Routes>
          <Route path="/" element={<ClassList />} />
          <Route path="/classes/:id" element={<ClassDetail />} />
          <Route path="/classes/add" element={<CreateClass />} />
          <Route
            path="/enrollments/classes/:id"
            element={<ClassParticipant />}
          />
          <Route
            path="/enrollments/participants/:id"
            element={<ParticipantClass />}
          />
          <Route path="/participants" element={<ParticipantList />} />
          <Route path="/participants/add" element={<CreateParticipant />} />
          <Route path="/participants/:id" element={<ParticipantDetail />} />
        </Routes>
      </div>
    </div>
  );
}

export default App;
