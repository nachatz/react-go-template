import React, { useState } from "react";
import Modal from "../../../components/Modal/Modal";

export default function Content() {
  const [opened, setOpened] = useState(false); 

  const openModal = () => {
    setOpened(true); 
  };

  return (
    <div className="mx-auto max-w-2xl lg:text-center mt-10">
      <h2 className="text-base font-semibold leading-7 text-indigo-600">
        Train Your Machine Learning Models
      </h2>
      <p className="mt-2 text-3xl font-bold tracking-tight text-gray-900 sm:text-4xl">
        Just upload your data, and we'll handle the rest
      </p>
      <p className="mt-6 text-lg leading-8 text-gray-600">
        MyAIMaker optimizes the procedure of training machine learning models by
        offering a streamlined user interface that facilitates the generation,
        engagement, and operationalization of state-of-the-art technology
      </p>
      <div className="flex justify-center mt-8">
        <button
          className="bg-primary-500 text-white font-semibold py-2 px-5 rounded-md shadow-md hover:bg-indigo-400 focus:outline-none focus:ring-1 focus:ring-indigo-400 focus:ring-offset-1"
          onClick={openModal} 
        >
          Try it now (it's free!)
        </button>
      </div>
      <Modal open={opened} setOpen={setOpened} />
    </div>
  );
}
