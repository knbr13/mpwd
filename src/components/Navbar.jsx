const Navbar = () => {
  return (
    <div className='navbar'>
      <span className="logo">iLine</span>
      <div className="user">
        <img src={""} alt="" />
        <span>username</span>
        <button>logout</button>
      </div>
    </div>
  )
}

export default Navbar