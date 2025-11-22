import { Link } from "react-router-dom";

export default function Navbar() {
  return (
    <nav className="bg-white shadow-md sticky top-0 z-50">
      <div className="max-w-7xl mx-auto px-6 h-16 flex items-center justify-between">
        <Link to="/" className="text-2xl font-bold text-blue-600">
          Skillhub
        </Link>
        <div className="flex gap-6 text-gray-700 font-medium">
          <Link className="hover:text-blue-600" to="/">
            Classes
          </Link>
          <Link className="hover:text-blue-600" to="/participants">
            Participants
          </Link>
        </div>
      </div>
    </nav>
  );
}
