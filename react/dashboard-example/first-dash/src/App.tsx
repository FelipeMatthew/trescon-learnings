import { useState } from 'react';
import Main from './components/Main';
import Sidebar from './components/Sidebar';
import './index.css';

function App() {
  const [isOpen, setIsOpen] = useState(true);
  
  const toggleSideBar = () => {
    setIsOpen(!isOpen)
  } 

  return (
    <div className='flex h-screen font-sans'>
      <Sidebar isOpen={isOpen} toggleSideBar={toggleSideBar}/>
      <Main/> 
    </div>
  );
}

export default App;
