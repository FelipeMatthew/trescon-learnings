import type React from "react";

interface Props {
  name: string;
  icon: string;
}

const CategoryItem: React.FC<Props> = ({ name, icon }) => {
  return (
    <div className="flex flex-col items-center">
      <div className=" p-4 rounded-xl hover:bg-gray-200 duration-300 cursor-pointer">
        <img src={icon} alt={name} className="w-10" />
      </div>

      <span className="font-medium text-gray-600">{name}</span>
    </div>
  );
};

export default CategoryItem;
