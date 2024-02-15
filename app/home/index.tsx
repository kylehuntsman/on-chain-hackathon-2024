'use client';
import HomeHeader from '../../src/pageComponents/home/HomeHeader';
import TransactionInput from './TransactionInput';

/**
 * Use the page component to wrap the components
 * that you want to render on the page.
 */

type ContainerProps = {
  amount: number;
  assetName: string;
  requestingUser: string;
  status: string;
}
function TransactionContainer({amount, assetName, requestingUser, status}: ContainerProps) {
  return <div>
    <div className='text-lg text-black p-2 text-center'>Transaction Details</div>
    <div className='text-lg text-black p-2 text-center'>Amount: {amount} {assetName}</div>
    <div className='text-lg text-black p-2 text-center'>From: {requestingUser}</div>
    <div className='text-lg text-black p-2 text-center'>To: You</div>
    <div className='text-lg text-black p-2 text-center'>Status: {status}</div>
  </div>
}

export default function HomePage() {
  const requestingUser = "User 1";

  const amount = 0.01;
  const assetName = "BTC";
  const status = "Pending"

  const title = `${requestingUser} has requested ${amount} ${assetName} from you!`;


  return (
    <>
    <HomeHeader />
    <main className="container mx-auto flex flex-col px-8 py-28 bg-black">
      <div className="container min-h-96 min-w-600px max-w-3/6">
        <div className='mx-auto justify-center flex text-3xl pb-4'>REQUEST YOUR CRYPTO</div>
        <div className="bg-white mx-8 rounded-md min-h-96 min-w-3/6">
          <div className="border-b-2">
            <div className='text-lg text-black p-2 text-center'>{title}</div>
          </div>
          <TransactionInput/>
          </div>
      </div>
    </main>
    </>
  );
}
