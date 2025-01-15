"use client";
import React from 'react';
import CartItem from '../components/CartItem';
import { useAuth } from '../context/AuthContext';
import Carts from '../lib/carts/Carts';
import { CartItemProps } from '../components/CartItem';
import products from '../lib/items/ProductList';

const Cart = () => {
    const { isAuthenticated, user } = useAuth();
    const cartProducts = Carts.find(c => c.customer_id === user?.userid)?.products || [];
    const productIDs = cartProducts.length == 0 ? [] : cartProducts.map(item => item.product_id);

    const productToCartItem = (product_id: string, quantity: number): CartItemProps | null => {

        if (!product_id) return null;

        const product = products.find(p => p.product_id === product_id);
        if (!product) return null;

        return {
            Item: {
                name: product.name,
                price: product.price,
                rating: product.rating,
                quantity: quantity,
                image: product.images?.[0] || "X"
            }
        };
    };



    return (
        <div className="container h-screen flex flex-row">
            {isAuthenticated ? (
                <>
                    <div className="relative w-3/5 flex flex-col">
                        <h1 className='font-bold text-2xl mt-4 ml-8'>Your Cart</h1>
                        {cartProducts.length > 0 ? (
                            cartProducts.map(p => {
                                const item = productToCartItem(p.product_id, p.quantity);
                                return item ? (
                                    <div key={p.product_id} className='mx-12 my-4'>
                                        <CartItem Item={item.Item} />
                                    </div>
                                ) : null;
                            })
                        ) : (
                            <p>Sepetiniz boş.</p>
                        )}
                    </div>
                    <div className="relative w-2/5 flex justify-center py-4 ">
                        <div className="w-4/5 h-1/5 flex mt-12 flex-col border-2 rounded-md shadow-lg justify-center px-12">
                            <h1 className='font-bold text-xl'>Total: $28998</h1>
                            <button className='bg-primary w-[90%] h-[36px] mb-[2px] rounded-full'>
                                Complate Order
                            </button>
                        </div>
                    </div>
                </>
            ) : (
                <h1>Not Found!</h1>
            )}
        </div>
    );
};

export default Cart;
