import { FaArrowUpLong, FaArrowDown } from "react-icons/fa6";
import { MdMoneyOff } from "react-icons/md";
import { FaMoneyBillTrendUp } from "react-icons/fa6";
import { FaChartBar } from "react-icons/fa6";
import {
  FaTachometerAlt,
  FaUser,
  FaCog,
  FaChartLine,
  FaDollarSign,
  FaUsers,
  FaFileAlt,
} from "react-icons/fa";

import drink from "../assets/drink.png";
import fruits from "../assets/fruits.png";
import travel from "../assets/travel.png";
import music from "../assets/music.png";
import joystick from "../assets/joystick.png";
import laptop from "../assets/laptop.png";
import mobile from "../assets/mobile.png";
import food from "../assets/food.png";

export const menuList = [
  { icon: <FaTachometerAlt className="text-xl" />, name: "Dashboard" },
  { icon: <FaUser className="text-xl" />, name: "Profile" },
  { icon: <FaCog className="text-xl" />, name: "Settings" },
  { icon: <FaChartLine className="text-xl" />, name: "Analytics" },
  { icon: <FaDollarSign className="text-xl" />, name: "Sales Report" },
  { icon: <FaUsers className="text-xl" />, name: "Customers" },
  { icon: <FaFileAlt className="text-xl" />, name: "Invoices" },
];

export const categories = [
  { name: "Drink", icon: drink },
  { name: "Fruits", icon: fruits },
  { name: "Travel", icon: travel },
  { name: "Music", icon: music },
  { name: "Joystick", icon: joystick },
  { name: "Laptop", icon: laptop },
  { name: "Mobile", icon: mobile },
  { name: "Food", icon: food },
  { name: "Drink", icon: drink },
  { name: "Fruits", icon: fruits },
  { name: "Travel", icon: travel },
  { name: "Music", icon: music },
  { name: "Joystick", icon: joystick },
  { name: "Laptop", icon: laptop },
  { name: "Mobile", icon: mobile },
  { name: "Food", icon: food },
];

export const cardData = [
  {
    title: "Average Revenue",
    revenue: "$129K",
    growth: {
      percentage: "0.5%",
      icon: FaArrowUpLong,
      color: "text-green-700",
    },
    image: FaMoneyBillTrendUp,
    bgColor: "bg-green-50",
    linkText: "Click now",
    linkHref: "#",
  },
  {
    title: "Total Sales",
    revenue: "$500K",
    growth: {
      percentage: "1.2%",
      icon: FaArrowUpLong,
      color: "text-green-700",
    },
    image: FaChartBar,
    bgColor: "bg-yellow-50",
    linkText: "View details",
    linkHref: "#",
  },
  {
    title: "Monthly Profit",
    revenue: "$75K",
    growth: {
      percentage: "-0.8%",
      icon: FaArrowDown,
      color: "text-red-700",
    },
    image: MdMoneyOff,
    bgColor: "bg-blue-50",
    linkText: "See report",
    linkHref: "#",
  },
];

// chart data
// const [series] = useState([
//   {
//     name: "series1",
//     data: [31, 40, 28, 51, 42, 109, 100],
//   },
//   {
//     name: "series2",
//     data: [11, 32, 45, 32, 34, 52, 41],
//   },
// ]);

// const [options] = useState({
//   chart: {
//     height: 350,
//     type: "area",
//   },
//   dataLabels: {
//     enabled: false,
//   },
//   stroke: {
//     curve: "smooth",
//   },
//   xaxis: {
//     type: "datetime",
//     categories: [
//       "2018-09-19T00:00:00.000Z",
//       "2018-09-19T01:30:00.000Z",
//       "2018-09-19T02:30:00.000Z",
//       "2018-09-19T03:30:00.000Z",
//       "2018-09-19T04:30:00.000Z",
//       "2018-09-19T05:30:00.000Z",
//       "2018-09-19T06:30:00.000Z",
//     ],
//   },
//   tooltip: {
//     x: {
//       format: "dd/MM/yy HH:mm",
//     },
//   },
// });
