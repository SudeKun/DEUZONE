import ItemCart from "./components/ItemCart";
import ProductSlider from "./components/ProductSlider";
import Slider from "./components/Slider";
import TopCategories from "./components/TopCategories";
import SliderImages from "./lib/images/SilderImages";

export default function Home() {
  return (
    <div className="min-h-screen flex flex-col items-center space-y-6">
      <Slider images={SliderImages}/>
      <TopCategories />
      <ProductSlider />
    </div>
  );
}
