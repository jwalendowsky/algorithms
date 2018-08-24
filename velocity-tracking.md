
# Velocity Tracking

## Chart

![alt text](velocity-tracking.jpg "Velocity Tracking Chart")

## Data

|Sprint|Engineers|Planned|Achieved|Achieved/Engineer|Achieved/Engineer(MA 3 Weeks)|
|:-|-:|-:|-:|-:|-:|velocityTrack => 
    velocityTrack.map(row => `[${row.milestone.title}](${row.milestone.url})|${row.milestone.capacity}|${row.statistics.commited}|${row.statistics.velocity}|${row.statistics.velocityPerEngineer.toFixed(DECIMALS)}||`)
      .reduce((lines, line) => lines + '\n' + line, '')
