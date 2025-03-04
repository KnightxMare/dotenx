import { ErrorMessage } from '@hookform/error-message'
import { InputHTMLAttributes } from 'react'
import { Control, Controller, FieldErrors } from 'react-hook-form'

interface FieldProps extends InputHTMLAttributes<HTMLInputElement> {
	label?: string
	errors: FieldErrors
	name: string
	// eslint-disable-next-line @typescript-eslint/no-explicit-any
	control: Control<any>
}

export function Field({ label, errors, control, ...rest }: FieldProps) {
	return (
		<div className="flex flex-col gap-1">
			{label && (
				<label htmlFor={rest.name} className="text-sm font-bold">
					{label}
				</label>
			)}
			<Controller
				control={control}
				name={rest.name}
				defaultValue="" // This is a bit tricky, maybe needs to change for different value types in future
				render={({ field: { onChange, value, name, ref } }) => (
					<input
						className="px-2 py-1 border rounded-lg border-slate-400 placeholder:text-slate-500 outline-rose-500 focus:ring-0 focus:border-slate-400 form-input focus:outline outline-2 outline-offset-[-1px]"
						id={rest.name}
						onChange={onChange}
						value={value}
						ref={ref}
						autoComplete="off"
						{...rest}
						name={name}
					/>
				)}
			/>

			{rest.name && <FieldError errors={errors} name={rest.name} />}
		</div>
	)
}

Field.displayName = 'Field'

interface FieldErrorProps {
	errors?: FieldErrors
	name: string
}

export function FieldError({ name, errors }: FieldErrorProps) {
	return (
		<ErrorMessage
			name={name}
			errors={errors}
			render={({ message }) => <p className="text-xs text-red-600">{message}</p>}
		/>
	)
}
