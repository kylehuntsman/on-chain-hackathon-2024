import React from 'react';

type TransactionDetailsProps = {
  amount: number;
  assetName: string;
  from: string;
  to: string;
}

function TransactionDetails({amount, assetName, from, to}: TransactionDetailsProps) {
  const transactionDetails = ``;

  return (
    <div>
      <div className='text-lg text-black p-2 text-center'>Transaction Details</div>
      <div className='text-lg text-black p-2 text-center'>Amount: {amount} {assetName}</div>
      <div className='text-lg text-black p-2 text-center'>From: {from}</div>
      <div className='text-lg text-black p-2 text-center'>To: {to}</div>
      <div className='text-lg text-black p-2 text-center'>
         <img src="path_to_your_image.jpg" alt="description_of_image" />
      </div>
      <div className='text-lg text-black p-2 text-center'>
        <button className='bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded'>
          Sign Transaction
        </button>
      </div>
    </div>
  );
}

export default TransactionDetails;