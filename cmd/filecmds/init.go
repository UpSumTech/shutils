package filecmds

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	parseShortDesc = `Prints examples of commands to handle file operations`
	parseLongDesc  = `Prints examples of commands to handle file operations`
	parseExample   = `
	### Available commands for file operations
	shutils file`
)

func Init() *cobra.Command {
	cmd := &cobra.Command{
		Use:              "file [no options!]",
		Short:            parseShortDesc,
		Long:             parseLongDesc,
		Example:          parseExample,
		TraverseChildren: true,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(`
# search for a string in the files of this dir
find . -type f -name '*.sh' -print | xargs grep 'string'

# List all directories
find . -type d -ls

# find and replace some text in a file in one line
find . -type f -name "*.json" -print | xargs grep -i 'string' | awk '{print $1}' | sed -e 's#:##g' | xargs -n 1 -I % sed -i -e 's#"string"#"newstring"#g' %

# find files with certain permission settings
find . -type f -perm 600 | ifne echo "executable files found"

# Find the actual number of cores on the machine. Single core might have a core id of 0.
cat /proc/cpuinfo | grep 'core id'

# find files with 600 permission settings across ssh dirs of 3 users in parallel
parallel -j3 -- "find /home/developer/.ssh -type f -perm 600" "find /root/.ssh -type f -perm 600" "find /home/ubuntu/.ssh -type f -perm 600"
egrep '(cal|date)' utils.sh # Find the strings in the file

# Fill some line numbers into a file
for i in {1..10}; do echo $i >> foo; done
# Use pee to pipe stdin to multiple files
combine foo or bar | pee 'sort -n | uniq >sorted' 'sort -nr | uniq >reverse_sorted'
# Sort the file numerically and add timestamps to the beginning of each line with sub-second resolution
cat sorted | ts -s "%Y/%m/%d:%H:%M:%.S" | sed -e 's#1970/01#2017/08#g;' | sponge sorted
sort -nr reverse_sorted | ts -s "%Y/%m/%d:%H:%M:%.S" | sed -e 's#1971/01#2017/08#g;' | sponge reverse_sorted
# Easily add line numbers and then add a tiume stamp between garbage text and the line numbers to a file
nl -bt -s " $(date +"%Y-%m-%d %H:%M:%S") " foo | sponge foo
# Try and generate something like a log file sythetically by adding a known prefix to the begining of a line
cat foo | awk -v prefix="[INFO] - 198.21.11.30" '{print prefix $0}' | sponge foo

# Applying a diff and patch for distributing changes quickly
diff -sub file1 file2 # shows diff between files file1 and file2 ignoring whitespaces
diff -Naur file1 file2 >v1 # generates a patch compatible diff between files file1 and file2 and puts it in patch v1
cat v1 | patch -p1 # apply patch v1 to file of choice interactively

shuf -i 1-100000 -n 500 > rand_numbers # Generates random numbers in the given range
cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 32 | head -n 1 # Generates random alphanumeric string that you can use for secrets or keys
ls /usr/bin | fold - -s -w 40 # Get the files in dir as a list

# Compress and decompress
tar -zcvf sorted.tar.gz sorted
tar -zcvf reverse_sorted.tar.gz reverse_sorted
zip -r sorted.zip sorted
zip -r reverse_sorted.zip reverse_sorted

# Join all bin folders and create a PATH like var
find ~ -name '*bin*' -type d | paste -d : -s -

# Change the ownership of all files belonging to a specfic user or group
find / -uid 1004 -exec chown -v 1010:1010 {} \;
find / -gid 1010 -exec chown -v 1001:1001 {} \;

readlink /usr/local/bin/awk # Quickly get where the file is pointing to

stat foo # Give detailed info about a file

# Tarpipe example to copy src to dest preserving perms and special flags etc
(cd src && tar -cf - .) | (cd dest && tar -xpf -)
			`)
		},
	}

	return cmd
}
