import type { NextPage } from "next";
import { useSelector } from "react-redux";
import Head from "next/head";
import Image from "next/image";

const Home: NextPage = () => {
  const val = useSelector((state) => state);
  console.log(val);
  return (
    <div className="h-screen w-screen overflow-auto bg-gray-800">
      <div className="h-full w-full flex flex-col space-y-10 items-center justify-center">
        <h1 className="text-white md:text-4xl text-3xl font-mono font-light">
          Systopher
        </h1>
        <h1 className="text-white md:text-7xl text-4xl">Coming Soon!</h1>
      </div>
    </div>
  );
};

export default Home;
