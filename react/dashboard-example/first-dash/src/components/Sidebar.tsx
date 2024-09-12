import { GiFireDash } from "react-icons/gi";
import { menuList } from "../constants";
import { FaTimes } from "react-icons/fa";
import { FaBars } from "react-icons/fa6";

interface Props {
  toggleSideBar: () => void;
  isOpen: boolean;
}

const Sidebar: React.FC<Props> = ({ toggleSideBar, isOpen }) => {
  return (
    <div
      className={`bg-gray-800 text-white fixed top-0 p-5 left-0 h-full flex flex-col justify-between py-5 ${
        isOpen ? "w-52" : "w-20"
      } duration-300`}
    >
      {/* Logo */}
      <div
        className={`${
          isOpen ? "" : "flex flex-col items-center justify-center"
        } duration-300`}
      >
        <div className="flex items-center justify-center gap-4">
          <GiFireDash className="text-2xl cursor-pointer hover:text-gray-400" />
          <span className={`text-2xl ${isOpen ? "" : "hidden"}`}>
            Felpola Dash
          </span>
        </div>
      </div>
      {/* Menu List */}
      <nav className="flex ">
        <ul className="flex flex-col gap-6 ">
          {menuList.map((item, index) => (
            <li
              key={index}
              className="flex text-xl items-center gap-3 cursor-pointer hover:text-gray-400"
            >
              <span>{item.icon}</span>
              <span className={`${isOpen ? "" : "hidden"} duration-300`}>
                {item.name}
              </span>
            </li>
          ))}
        </ul>
      </nav>

      <div
        className="flex items-center justify-center p-3 cursor-pointer hover:bg-gray-700 rounded"
        onClick={toggleSideBar}
      >
        {/* Toggle */}

        {isOpen ? <FaTimes /> : <FaBars />}
      </div>
    </div>
  );
};

export default Sidebar;
