import ApexChart from "./ApexChart";
import Categories from "./Categories";
import Header from "./Header";

const Main = () => {
  return (
    <div className="flex-1">
      <Header />
      <Categories />

      <div className="flex gap-4 p-10">
        <ApexChart />
      </div>
    </div>
  );
};

export default Main;
