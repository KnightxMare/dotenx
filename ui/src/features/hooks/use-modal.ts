import { useAtom } from 'jotai'
import { modalAtom } from '../atoms'

export enum Modals {
	NodeSettings = 'node-settings',
	EdgeSettings = 'edge-settings',
	TriggerSettings = 'trigger-settings',
	SaveAutomation = 'save-automation',
	TaskLog = 'task-log',
	NewIntegration = 'new-integration',
	NewTrigger = 'new-trigger',
	HotKeys = 'hot-keys',
	AutomationYaml = 'automation-yaml',
	DeleteAutomation = 'delete-automation',
	NewProvider = 'new-provider',
	NewProject = 'new-project',
	NewTable = 'new-table',
	NewColumn = 'new-column',
	TableEndpoints = 'table-endpoints',
	InteractionResponse = 'table-endpoints',
	TemplateEndpoint = 'template-endpoint',
	UserManagementEndpoint = 'user-management-endpoint',
	InteractionBody = 'interaction-body',
}

export function useModal() {
	const [modal, setModal] = useAtom(modalAtom)

	return {
		...modal,
		open: (kind: Modals, data: unknown = null) => setModal({ isOpen: true, kind, data }),
		close: () => setModal({ isOpen: false, kind: null, data: null }),
	}
}
