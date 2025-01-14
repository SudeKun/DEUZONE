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
                    <div className="relative w-4/5 flex flex-col">
                        <h1>Your Cart</h1>
                        {cartProducts.length > 0 ? (
                            cartProducts.map(p => {
                                const item = productToCartItem(p.product_id, p.quantity);
                                return item ? (
                                    <div key={p.product_id} className='m-4'>
                                        <CartItem Item={item.Item} />
                                    </div>
                                ) : null;
                            })
                        ) : (
                            <p>Sepetiniz boş.</p>
                        )}
                    </div>
                    <div className="relative w-1/5 flex items-center">
                        
                    </div>
                </>
            ) : (
                <h1>Not Found!</h1>
            )}
        </div>
    );
};

export default Cart;
