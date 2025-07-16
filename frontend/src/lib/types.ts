import { Status } from '$lib/enums/status';

export type Todo = {
	id: string;
	text: string;
	status: Status;
};
