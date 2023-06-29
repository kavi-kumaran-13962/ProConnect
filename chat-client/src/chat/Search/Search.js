import React, { useState } from 'react';
import "./Search.css"
const Search = () => {
    const [searchTxt, setsearchTxt] = useState('');

    const handleInputChange = (e) => {
        setsearchTxt(e.target.value);
    };

    const handleSend = () => {
        // Handle sending the searchTxt here
        console.log(searchTxt);
        setsearchTxt('');
    };

    return (
        <div className='search'>
            <input 
                type="text" 
                value={searchTxt} 
                onChange={handleInputChange} 
                placeholder="search..."     
                className='search__inp'
            />
            <div 
                onClick={handleSend} 
                className='search__sendDiv'
            >
                <i className='search__sendicon fa fa-search icon'></i>
            </div>
        </div>
    );
};

export default Search;
