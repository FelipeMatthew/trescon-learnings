import { useState } from "react";
import Main from "./components/Main";
import Sidebar from "./components/Sidebar";
import "./index.css";

function App() {
  const [isOpen, setIsOpen] = useState(true);

  const toggleSideBar = () => {
    setIsOpen(!isOpen);
  };

  return (
    <div className="flex h-screen font-sans">
      <Sidebar isOpen={isOpen} toggleSideBar={toggleSideBar} />
      <div
        className={`flex-1 transition-all duration-300 ${
          isOpen ? "ml-52" : "ml-20"
        }`}
      >
        <Main />
      </div>
    </div>
  );
}

export default App;
