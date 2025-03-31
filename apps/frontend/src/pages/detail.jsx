import React, { useState, useEffect } from 'react';
import { useParams, Link } from 'react-router-dom';

const UserDetail = () => {
  const { id } = useParams(); // If there's an ID, we're editing an existing user.

  // Initialize form data with empty strings.
  const [formData, setFormData] = useState({
	hash: '',
    user_name: '',
    first_name: '',
    last_name: '',
    email: '',
    department: '',
  });

  // If editing, fetch the existing user data.
  useEffect(() => {
    if (id) {
      fetch(`/api/user/${id}`)
        .then((res) => {
          if (!res.ok) throw new Error('Error fetching user data');
          return res.json();
        })
        .then((data) => {
          setFormData({
            user_name: data.user_name || '',
            first_name: data.first_name || '',
            last_name: data.last_name || '',
            email: data.email || '',
            user_status: data.user_status || '',
            department: data.department || '',
          });
        })
        .catch((err) => console.error('Error:', err));
    }
  }, [id]);

  // Handle input changes.
  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData((prev) => ({ ...prev, [name]: value }));
  };

  // Handle form submission.
  const handleSubmit = (e) => {
    e.preventDefault();
    const url = id ? `/api/user/${id}` : `/api/user`;
    const method = id ? 'PUT' : 'POST';

    fetch(url, {
      method,
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(formData),
    })
      .then((res) => {
        if (!res.ok) throw new Error('Error saving user');
        return res.json();
      })
      .catch((err) => console.error('Error:', err));
  };

  return (
    <main className="max-w-xl mx-auto p-4">
      <h2 className="text-2xl font-bold mb-4">
        {id ? 'Edit User' : 'Create User'}
      </h2>
      <form onSubmit={handleSubmit} className="space-y-4">
        <div>
          <label className="block mb-1 font-semibold">Username</label>
          <input
            type="text"
            name="user_name"
            value={formData.user_name}
            onChange={handleChange}
            className="w-full border border-gray-300 rounded p-2"
            required
          />
        </div>
        <div>
          <label className="block mb-1 font-semibold">First Name</label>
          <input
            type="text"
            name="first_name"
            value={formData.first_name}
            onChange={handleChange}
            className="w-full border border-gray-300 rounded p-2"
            required
          />
        </div>
        <div>
          <label className="block mb-1 font-semibold">Last Name</label>
          <input
            type="text"
            name="last_name"
            value={formData.last_name}
            onChange={handleChange}
            className="w-full border border-gray-300 rounded p-2"
            required
          />
        </div>
        <div>
          <label className="block mb-1 font-semibold">Email</label>
          <input
            type="email"
            name="email"
            value={formData.email}
            onChange={handleChange}
            className="w-full border border-gray-300 rounded p-2"
            required
          />
        </div>
        <div>
          <label className="block mb-1 font-semibold">Department</label>
          <select
            name="department"
            value={formData.department}
            onChange={handleChange}
            className="w-full border border-gray-300 rounded p-2"
          >
            <option value="">Select a Department</option>
            <option value="accounting">Accounting</option>
            <option value="it">Information Technology</option>
            <option value="hr">Human Resources</option>
            <option value="marketing">Marketing</option>
            <option value="sales">Sales</option>
            <option value="engineering">Engineering</option>
            <option value="operations">Operations</option>
          </select>
        </div>
        <button
          type="submit"
          className="bg-blue-500 hover:bg-blue-600 text-white px-4 py-2 rounded"
        >
          {id ? 'Update User' : 'Create User'}
        </button>
		<Link className="bg-white-500 hover:bg-blue-600 text-black px-4 py-2 rounded" to="/">
			Back
		</Link>
      </form>
    </main>
  );
};

export default UserDetail;
