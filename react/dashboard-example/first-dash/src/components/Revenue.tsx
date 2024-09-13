import { cardData } from "../constants";
import Title from "../ui/Title";
import RevenueCard from "./RevenueCard";

const Revenue = () => {
  return (
    <div className="w-[350px]">
      <Title>Revenue</Title>

      <div className="flex flex-col gap-4">
        {cardData.map((card, index) => (
          <div key={index}>
            <RevenueCard card={card}/>
          </div>
        ))}
      </div>
    </div>
  );
};

export default Revenue;
