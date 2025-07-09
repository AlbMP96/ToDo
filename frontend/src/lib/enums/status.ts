export enum Status {
	Pending = 'pending',
	InProgress = 'in_progress',
	Completed = 'completed'
}

export const statusLabels: Record<Status, string> = {
	[Status.Pending]: 'Pendiente',
	[Status.InProgress]: 'En progreso',
	[Status.Completed]: 'Completado'
};
