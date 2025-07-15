import { Button } from '@/components/ui/button'
import { Calendar } from '@/components/ui/calendar'
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from '@/components/ui/popover'
import { format } from 'date-fns'

interface Props {
  date: Date | undefined
  onSelectDate: (d: Date | undefined) => void
}

const DatePicker = ({ date, onSelectDate }: Props) => {
  return (
    <Popover>
      <PopoverTrigger asChild>
        <Button
          variant="outline"
          data-empty={!date}
          className="data-[empty=true]:text-muted-foreground w-full justify-start text-left font-normal"
        >
          {date ? format(date, "PPP") : <span>Choose date</span>}
        </Button>
      </PopoverTrigger>
      <PopoverContent className="w-auto p-0">
        <Calendar mode="single" selected={date} onSelect={onSelectDate} />
      </PopoverContent>
    </Popover>
  )
}

export default DatePicker