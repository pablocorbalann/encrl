# This small script has been created to automate the task of 
# changing the parche for the Encrl version.
#
# If the version is something like:
#   version.number.parche
#
# For example
#   1.0.3
#
# This script modifies the 'parche' section to `parche=parche+1`
#
# The script works in the file version.txt by default
# Modify the name of the file here:
FILE_ROUTE="version.txt"
VERSION_INCREMENT=1

# DO NOT TOUCH ANYTHING FROM HERE TO THE END OF THE CODE
# --------------------------------------------------------------------
# Before making any important change inside the version file we have
# to work with the file routes and handle the errors.
#   - Convert from the relative to the absolute path using thr readlink
#     command, then store the value in a variable
#   - Read the version parche
#   - Increase it by 1 (nv(parche)=v+1||v++)
FILE=$(readlink ${FILE_ROUTE} -f)
PARCHE=$(sed -n "3p" ${FILE})
PARCHE=$((PARCHE+VERSION_INCREMENT))
echo $PARCHE
# Once we have all this information, we have to actually modify the 
# file to increase the version. 
#   - Delete all the content inside the file
#   - Open the file itself
#   - Set the parche to line 3
sed -i "3s/.*/${PARCHE}/" ${FILE}
