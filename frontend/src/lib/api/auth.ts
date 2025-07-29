import { apiRoutes } from '$lib/config/apiRoutes';

export async function login(email: string, password: string) {
	const response = await fetch(apiRoutes.login, {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json'
		},
		body: JSON.stringify({ email, password }),
		credentials: 'include'
	});

	if (!response.ok) {
		throw new Error(`Error al iniciar sesión: ${response.statusText}`);
	}

	return response.json();
}

export async function create(email: string, password: string) {
	const response = await fetch(apiRoutes.register, {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json'
		},
		body: JSON.stringify({ email, password })
	});

	if (!response.ok) {
		throw new Error(`Error al crear cuenta: ${response.statusText}`);
	}

	return response.json();
}

export async function logout() {
	const response = await fetch(apiRoutes.logout, {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json'
		}
	});

	if (!response.ok) {
		throw new Error(`Error al cerrar sesión: ${response.statusText}`);
	}

	return response.json();
}
