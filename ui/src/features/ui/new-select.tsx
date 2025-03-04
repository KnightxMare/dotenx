import clsx from 'clsx'
import { useRef, useState } from 'react'
import { Control, Controller, FieldErrors } from 'react-hook-form'
import { IoChevronDown } from 'react-icons/io5'
import { useOutsideClick } from '../hooks'
import { Fade } from './animation/fade'
import { FieldError } from './field'
import { Loader } from './loader'

interface Option {
	label: string
	value: string
}

export interface NewSelectProps {
	label?: string
	errors?: FieldErrors
	name: string
	// eslint-disable-next-line @typescript-eslint/no-explicit-any
	control: Control<any>
	options?: Option[]
	placeholder?: string
	loading?: boolean
}

export function NewSelect({
	label,
	errors,
	name,
	control,
	options = [],
	placeholder,
	loading,
	...rest
}: NewSelectProps) {
	return (
		<div className="flex flex-col gap-1" {...rest}>
			<label htmlFor={name} className="text-sm font-bold">
				{label}
			</label>
			<SelectController
				name={name}
				control={control}
				options={options}
				placeholder={placeholder}
				loading={loading}
			/>
			{name && errors && <FieldError errors={errors} name={name} />}
		</div>
	)
}

interface SelectControllerProps {
	name: string
	control: Control
	options: Option[]
	placeholder?: string
	loading?: boolean
}

function SelectController({ control, name, options, placeholder, loading }: SelectControllerProps) {
	return (
		<Controller
			control={control}
			name={name}
			defaultValue=""
			render={({ field: { onChange, value } }) => {
				return (
					<RawSelect
						onChange={(newValue) => onChange((newValue as Option)?.value)}
						value={options.find((option) => option.value === value)}
						options={options}
						placeholder={placeholder}
						loading={loading}
					/>
				)
			}}
		/>
	)
}

function RawSelect({
	onChange,
	options,
	value,
	placeholder,
	loading,
}: {
	value?: Option
	onChange: (value: Option) => void
	options: Option[]
	placeholder?: string
	loading?: boolean
}) {
	const [isOpen, setIsOpen] = useState(false)
	const wrapperRef = useRef<HTMLDivElement>(null)
	useOutsideClick(wrapperRef, () => setIsOpen(false))

	return (
		<div className="relative space-y-1" ref={wrapperRef}>
			<button
				type="button"
				className={clsx(
					'flex items-center justify-between w-full px-2 py-1 border rounded-lg cursor-pointer border-slate-400 outline-rose-500',
					isOpen && 'outline outline-2 outline-offset-[-1px]'
				)}
				onClick={() => setIsOpen((isOpen) => !isOpen)}
			>
				{!value && <span className="text-slate-500">{placeholder}</span>}
				{value && <span>{value.label}</span>}
				<IoChevronDown className="text-slate-400" />
			</button>
			<div className="absolute inset-x-0 z-10">
				<Fade isOpen={isOpen}>
					<div className="p-1 bg-white border rounded-lg shadow-md border-slate-300">
						{loading && <Loader className="py-4" />}
						{!loading && options.length === 0 && (
							<div className="p-1.5 text-xs font-thin text-center">No options</div>
						)}
						{options.map((option) => (
							<button
								type="button"
								className="w-full px-2 py-1 text-left transition rounded-md hover:bg-rose-50 focus:bg-rose-50 outline-rose-500"
								key={option.label}
								onClick={() => {
									onChange(option)
									setIsOpen(false)
								}}
							>
								{option.label}
							</button>
						))}
					</div>
				</Fade>
			</div>
		</div>
	)
}
