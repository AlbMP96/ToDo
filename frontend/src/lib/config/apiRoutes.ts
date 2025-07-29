const baseUrl = 'http://localhost:8080';
const apiUrl = `${baseUrl}/api`;
const authUrl = `${baseUrl}/auth`;

export const apiRoutes = {
	base: baseUrl,
	login: `${authUrl}/login`,
	register: `${authUrl}/register`,
	logout: `${authUrl}/logout`
};
