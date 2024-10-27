import React, { useState } from 'react';
import axios from 'axios';
import './Home.css'; // Make sure to import the CSS file

function Home() {
    const [inputText, setInputText] = useState("");
    const [data, setData] = useState("");
    const [loading, setLoading] = useState(false);

    function handleOnChange(e) {
        setInputText(e.target.value);
    }

    function handleOnClear() {
        setInputText("");
        setData("");
    }

    function handleOnSubmit() {
      if(inputText == '') return;
        setLoading(true); 
        axios.post('http://localhost:3000/shorten', {
            url: inputText,
        })
        .then((response) => {
            setData(response.data.short_url);
        })
        .catch((error) => {
            console.error('Error:', error);
        })
        .finally(() => {
            setLoading(false); 
        });
    }

    return (
        <div className="flex items-center justify-center min-h-screen bg-gradient-to-r from-purple-400 via-pink-500 to-red-500 bg-fixed animate-gradient">
            <div className="bg-white rounded-lg shadow-lg p-8 w-full max-w-md">
                <h1 className="text-2xl font-bold text-center mb-6">URL Shortener</h1>
                <input
                    type="text"
                    placeholder="Enter URL here"
                    onChange={handleOnChange}
                    value={inputText} // Controlled input
                    className="w-full p-3 border border-gray-300 rounded mb-4 focus:outline-none focus:ring-2 focus:ring-blue-500"
                />
                <button
                    onClick={handleOnSubmit}
                    className="w-full bg-blue-500 text-white py-2 rounded hover:bg-blue-600 transition duration-200"
                >
                    {loading ? 'Shortening...' : 'Shorten URL'}
                </button>
                {data && (
                    <div className="mt-6 text-center">
                        <span className="block text-gray-700">Shortened URL:</span>
                        <span className="font-semibold text-blue-600">{data}</span>
                        <div className="mt-2">
                            <a 
                                href={`http://localhost:3000/redirect/${data}`} 
                                target="_blank" 
                                rel="noopener noreferrer" 
                                className="text-blue-500 underline"
                            >
                                Click Here
                            </a>
                            <button 
                                onClick={handleOnClear}
                                className="ml-4 text-red-500 hover:text-red-600"
                            >
                                Clear
                            </button>
                        </div>
                    </div>
                )}
            </div>
        </div>
    );
}

export default Home;
