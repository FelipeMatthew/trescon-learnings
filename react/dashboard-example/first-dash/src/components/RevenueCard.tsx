import type React from "react";
import { IconType } from "react-icons";

interface Props {
  card: RevenueCard;
}

interface Growth {
  percentage: string;
  icon: IconType;
  color: string;
}

interface RevenueCard {
  title: string;
  revenue: string;
  growth: Growth;
  image: IconType;
  bgColor: string;
  linkText: string;
  linkHref: string;
}

const RevenueCard: React.FC<Props> = ({ card }) => {
  return (
    <div>
      <card.image />

      <div>
        <span>{card.title}</span>

        <div>
          <h1>{card.revenue}</h1>
          <span>
            <card.growth.icon /> {card.growth.percentage}
          </span>
        </div>
      </div>
    </div>
  );
};

export default RevenueCard;
