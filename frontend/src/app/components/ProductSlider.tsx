'use client'
import React from 'react'
import ProductList from '../lib/items/ProductList'
import ItemCart from './ItemCart'
import { useState } from 'react'

const ProductSlider: React.FC = () => {

    const [currentIndex, setCurrentIndex] = useState(0);
    const itemsPerPage = 5;
    const visibleProducts = ProductList.slice(currentIndex, currentIndex + itemsPerPage);

    const goToNext = () => {
        if (currentIndex < ProductList.length - itemsPerPage) {
            setCurrentIndex(currentIndex + 1);
        }
    };

    const goToPrev = () => {
        if (currentIndex > 0) {
            setCurrentIndex(currentIndex - 1);
        }
    };


    return (
        <div className='relative flex flex-col w-screen'>
            <h1>DISCOUNT ITEMS</h1>
            <div className="flex flex-row w-full justify-center items-center">
                <button
                    onClick={goToPrev}
                    className="absolute left-10 top-1/2 transform -translate-y-1/2 text-red p-2 rounded-full"
                >
                    &lt;
                </button>
                <div className="flex w-10/12 justify-between overflow-x-auto py-4">
                    {
                        visibleProducts.map((Product) => (
                            <div key={Product.product_id}>
                                <ItemCart
                                    title={Product.name}
                                    price={Product.price}
                                    image={Product.images[0]}
                                    stars={Product.rating}
                                    comments={Product.review_count}
                                />
                            </div>
                        ))
                    }
                </div>
                <button
                    onClick={goToNext}
                    className="absolute right-10 top-1/2 transform -translate-y-1/2 text-red p-2 rounded-full"
                >
                    &gt;
                </button>
            </div>
        </div>
    )
}

export default ProductSlider;
