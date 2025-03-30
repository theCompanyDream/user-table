import React, { useContext, useEffect } from 'react';
import {UserContext, Table} from '../components'

const UserTable = () => {
	const { users, setUsers } = useContext(UserContext);

	useEffect(() => {
		if (users.users.length == 0) {
			fetch('/api/users')
			.then(response => response.json())
			.then(data => {
				setUsers(data);
			})
		}
	}, [setUsers, users]);
	return (
		<main>
			{users &&
				<Table users={users.users} page={users.page} totalPages={users.Length}  />
			}
		</main>
	);
}

export default UserTable;