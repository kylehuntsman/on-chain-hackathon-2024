import React, { useState } from 'react';

function TransactionInput() {
  const [chain, setChain] = useState('');
  const [amount, setAmount] = useState('');
  const [toAddress, setToAddress] = useState('');
  const [message, setMessage] = useState('');
  const [link, setLink] = useState('');

  const handleSubmit = (event) => {
    event.preventDefault();
    // Handle form submission here
  };

  return (
    <form onSubmit={handleSubmit} className="mx-auto flex flex-col justify-center">
      <div className="block text-black p-4">
        Chain:
        <input className="w-3/6 h-10 border-2 border-gray-300 p-2 bg-white" type="text" value={chain} onChange={(e) => setChain(e.target.value)} />
      </div>
      <div className="block text-black p-4">
        Amount:
        <input className="w-3/6 h-10 border-2 border-gray-300 p-2 bg-white" type="text" value={amount} onChange={(e) => setAmount(e.target.value)} />
      </div>
      <div className="block text-black p-4">
        To Address:
        <input className="w-3/6 h-10 border-2 border-gray-300 p-2 bg-white" type="text" value={toAddress} onChange={(e) => setToAddress(e.target.value)} />
      </div>
      <div className="block text-black p-4">
        Message:
        <input className="w-3/6 h-10 border-2 border-gray-300 p-2 bg-white" type="text" value={message} onChange={(e) => setMessage(e.target.value)} />
      </div>
      <div className="block text-black p-4">
        Link:
        <input className="w-3/6 h-10 border-2 border-gray-300 p-2 bg-white" type="text" value={link} onChange={(e) => setLink(e.target.value)} />
      </div>
      <input className="w-3/6 h-10 bg-blue-500 text-white p-2 cursor-pointer m-4 border-rad" type="submit" value="Submit" />
    </form>
  );
}

export default TransactionInput;