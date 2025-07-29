import { jwtVerify } from 'jose';
import '$lib/server/env';

const secretString = process.env.JWT_SECRET_KEY;
if (!secretString) {
	throw new Error('JWT_SECRET_KEY no está definida');
}
const encoder = new TextEncoder();
const secret = encoder.encode(secretString);

export async function verifyToken(token: string) {
	console.log('Secret Key:', secret);
	try {
		const { payload } = await jwtVerify(token, secret);
		return payload;
	} catch (error) {
		console.error('Token verification failed:', error);
		throw new Error('Invalid token');
	}
}
