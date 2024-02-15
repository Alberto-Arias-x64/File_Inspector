import { DateTime } from 'luxon'
import { format, formatISO } from 'date-fns'

const today = new Date()

console.log(DateTime.now().toString())
console.log(formatISO(today))
console.log(today.toISOString())
console.log(DateTime.now().toFormat('yyyy - MM - dd / hh:mm:ss a'))
console.log(format(today, 'yyyy - MM - dd / hh:mm:ss a'))