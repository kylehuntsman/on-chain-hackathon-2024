import React, { useEffect, useState } from 'react';
import { useRouter } from 'next/router';


type Transaction = {
  transaction: {
    address: string;
    amount: number;
  },
  qrCodeStr: string;
}

function RequestsPage() {

  const router = useRouter()
  const { id } = router.query

  // Make request to the server to get the transaction details

  // GET request to server
  const [resp, setResp] = useState<Transaction>({
    transaction: {
      address: '',
      amount: 0
    },
    qrCodeStr: ''
  });

  // useEffect(() => {
  //   async function getTransactionDetails(transactionId: string): Promise<Transaction> { 
  //     // Make request to the server to get the transaction details
  
  //     // GET request to server
  //     const response = await fetch(`${process.env.NEXT_PUBLIC_API_HOST}/transactions/${transactionId}`);
  //     const data = await response.json() as Transaction;
  
  //     return data;
  //   }
  //   getTransactionDetails(id as string)
  //     .then((res) => setResp(res))
  //     .catch((err) => console.log(err));
  // }, [id]);

  return (
    <div>
      <div className='bg-white'>
      <div className='text-lg text-black p-2 text-center'>Transaction Details</div>
      <div className='text-lg text-black p-2 text-center'>Amount: {resp.transaction.amount} ETH</div>
      <div className='text-lg text-black p-2 text-center'>To: {resp.transaction.address}</div>
      <div className='text-lg text-black p-2 text-center'>
         <img src={`${resp.qrCodeStr}`} alt="description_of_image" />
      </div>
      <div className='text-lg text-black p-2 text-center'>
        <button className='bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded'>
          Sign Transaction
        </button>
      </div>
      </div>
    </div>
  );
}

export default RequestsPage;