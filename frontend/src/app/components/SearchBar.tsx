import React from 'react'

const SearchBar = () => {
  return (
    <form className='relative w-[700px] h-[40px]'>
      <input type="text" placeholder='Type Here...' className='relative rounded-full w-[100%] h-[100%] pl-4 pr-24'/>
      <button type="submit" className='bg-white border-s text-textColor absolute w-[18%] h-[100%] rounded-r-full right-0 top-0' >Search</button>
    </form>
  )
}

export default SearchBar;
