'use client'

import { useState } from "react";
import Image, { StaticImageData } from "next/image";

interface SliderProps {
  images: StaticImageData[];
}

const Slider: React.FC<SliderProps> = ({ images }) => {
  const [currentIndex, setCurrentIndex] = useState(0);

  const handlePrev = () => {
    setCurrentIndex((prevIndex) =>
      prevIndex === 0 ? images.length - 1 : prevIndex - 1
    );
  };

  const handleNext = () => {
    setCurrentIndex((prevIndex) =>
      prevIndex === images.length - 1 ? 0 : prevIndex + 1
    );
  };

  return (
    <div className="relative flex w-screen h-[600px] overflow-hidden rounded-md justify-center mt-10">
      <div className="overflow-hidden w-11/12 h-full rounded-3xl border-2 border-borderprimary">
        <Image
          src={images[currentIndex]}
          alt={`Slide ${currentIndex}`}
          className="w-full h-full object-cover transition-transform duration-500"
        />

        <button
          onClick={handlePrev}
          className="absolute top-1/2 left-20 transform text-red p-2 rounded-full hover:bg-gray-700"
        >
          &#10094;
        </button>

        <button
          onClick={handleNext}
          className="absolute top-1/2 right-20 transform text-red p-2 rounded-full hover:bg-gray-700"
        >
          &#10095;
        </button>
      </div>
    </div>
  );
};

export default Slider;
