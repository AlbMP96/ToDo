import { redirect } from '@sveltejs/kit';
import { verifyToken } from '$lib/server/auth/verify';
import '$lib/server/env';

export const load = async ({ cookies, url }) => {
	const token = cookies.get('jwt') || '';
	console.log('Token recibido:', token);

	const publicRoutes = ['/login', '/register'];

	if (publicRoutes.includes(url.pathname)) {
		return {};
	}

	if (!token) {
		console.log('No hay token, redirigiendo...');
		throw redirect(302, '/login');
	}

	try {
		console.log('Verificando token...');
		const user = await verifyToken(token);
		console.log('Usuario verificado:', user);
		return { user };
	} catch (e) {
		console.error('Fallo al verificar token:', e);
		throw redirect(302, '/login');
	}
};
