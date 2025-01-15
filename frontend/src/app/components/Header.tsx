'use client'
import React from 'react'
import SearchBar from './SearchBar';
import Link from 'next/link';
import { useAuth } from '../context/AuthContext';
import { useEffect } from 'react';
import { useRouter } from 'next/navigation';



const Header = () => {

  const { isAuthenticated, user, logout } = useAuth();
  const router = useRouter();

  const handleClickProfile = () => {
    router.push(`/profile/${user?.userid}`);
  }

  return (
    <div className='relative h-[150px]'>
      <div className='bg-primary flex flex-row justify-between items-center h-[60%] px-8'>
        <h1 className='text-3xl font-bold text-textColor'>DEUZONE</h1>
        <SearchBar />
        <div className='w-[25%] flex justify-between items-center'>
          {
            isAuthenticated ? (
              <>
                <button onClick={handleClickProfile} className='bg-white w-[150px] h-[40px] rounded-md text-textColor'>Welcome, {user?.name}</button>
                <Link href='/cart'>
                  <button className='bg-white w-[150px] h-[40px] rounded-full text-textColor'>Cart</button>
                </Link>
                <button onClick={logout} className='bg-white w-[150px] h-[40px] rounded-full text-textColor'>Log Out</button>
              </>
            ) : (
              <Link href='/login'>
                <button className='bg-white w-[150px] h-[40px] rounded-full text-textColor'>Log In</button>
              </Link>
            )
          }
        </div>
      </div>
      <div className='bg-primary border-t flex flex-row justify-between items-center px-8 h-[40%]'>
        <div className='flex flex-row w-[15%] justify-between'>
          <h4 className='text-textColor'>All Categories</h4>
          <h4 className='text-textColor'>Help Center</h4>
        </div>
      </div>
    </div>
  )
}

export default Header;
