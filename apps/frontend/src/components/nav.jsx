import React, { useState, memo } from 'react';

const Navigation = memo(() => {
  const [isOpen, setIsOpen] = useState(false);

  return (
    <header className="bg-gray-800 text-white shadow w-full">
      <div className="max-w-7xl mx-auto px-4 py-3 flex items-center justify-between">
        {/* Logo or Title */}
        <div className="text-xl font-bold">User Administrator</div>
        {/* Desktop Navigation */}
        <nav className="hidden md:flex space-x-4">
          <a href="/about" className="hover:text-gray-300">Create</a>
          <a href="/services" className="hover:text-gray-300">Docs</a>
          <a href="/contact" className="hover:text-gray-300">Contact</a>
        </nav>
        {/* Mobile Menu Button */}
        <div className="md:hidden">
          <button
            onClick={() => setIsOpen(!isOpen)}
            className="focus:outline-none"
          >
            {isOpen ? (
              // Close (X) icon
              <svg
                className="w-6 h-6"
                fill="none"
                stroke="currentColor"
                strokeWidth="2"
                viewBox="0 0 24 24"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path strokeLinecap="round" strokeLinejoin="round" d="M6 18L18 6M6 6l12 12" />
              </svg>
            ) : (
              // Hamburger icon
              <svg
                className="w-6 h-6"
                fill="none"
                stroke="currentColor"
                strokeWidth="2"
                viewBox="0 0 24 24"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path strokeLinecap="round" strokeLinejoin="round" d="M4 6h16M4 12h16M4 18h16" />
              </svg>
            )}
          </button>
        </div>
      </div>
      {/* Mobile Navigation */}
      {isOpen && (
        <div className="md:hidden bg-gray-700">
          <nav className="px-4 py-2 space-y-1">
            <a href="/about" className="block px-3 py-2 rounded-md text-base font-medium hover:bg-gray-600">
              Create
            </a>
            <a href="/services" className="block px-3 py-2 rounded-md text-base font-medium hover:bg-gray-600">
              Docs
            </a>
            <a href="/contact" className="block px-3 py-2 rounded-md text-base font-medium hover:bg-gray-600">
              Contact
            </a>
          </nav>
        </div>
      )}
    </header>
  );
});

export default Navigation;
