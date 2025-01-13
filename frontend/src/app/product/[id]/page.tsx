"use client"

import { useParams } from 'next/navigation';
import products from '@/app/lib/items/ProductList';


const ProductPage = () => {
    const params = useParams(); 
    const id = params?.id;  // TypeScript için tip belirtildi
    const product = products.find((p) => p.product_id === id);

    /*
    useEffect(() => {
        if (id) {
            fetch(`/api/products/${id}`)
                .then((res) => res.json())
                .then((data) => setProduct(data));
        }
    }, [id]);
    */


    if (!product) return <p>Loading...</p>;

    return (
        <div className="container mx-auto p-8">
            <div className="flex">
                <div className="w-1/2">
                    {product.images.map((img, index) => (
                        <img key={index} src={img} alt={product.name} className="mb-4" />
                    ))}
                </div>
                <div className="w-1/2 flex flex-col justify-start p-4">
                    <h1 className="text-2xl font-bold">{product.name}</h1>
                    <p>{product.review_count} reviews - {product.rating}/5</p>
                    <p className="text-xl mt-2">Price: ${product.price}</p>
                    <button className="bg-blue-500 text-white px-4 py-2 mt-4">Add to Cart</button>
                </div>
            </div>

            
            <div className="mt-8">
                <h2 className="text-xl font-semibold">Description</h2>
                <p>{product.description}</p>
                <h3 className="text-lg font-semibold mt-4">Key Features</h3>
                <ul>
                    {product.features.map((feature, index) => (
                        <li key={index}>{feature}</li>
                    ))}
                </ul>
            </div>

        </div>
    );
};

export default ProductPage;
