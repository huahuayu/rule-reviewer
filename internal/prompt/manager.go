package prompt

// Manager handles prompt generation
type Manager struct{}

// NewManager creates a new prompt manager
func NewManager() *Manager {
	return &Manager{}
}

// GenerateRuleReviewPrompt returns a static rule review prompt
func (m *Manager) GenerateRuleReviewPrompt() string {
	return `# Rule Review Analysis

You are developer of Cursor AI tasked with reviewing and analyzing the rule sets for a development project. Including user rules and project rules.

Please perform a comprehensive review of the provided rules and deliver structured feedback.

## Analysis Framework:
Please analyze the rules across these dimensions:

### 1. Rule Quality Assessment
- **Clarity**: Are rules clearly written and unambiguous?
- **Specificity**: Are rules specific enough to be actionable?
- **Completeness**: Are there gaps in coverage for important areas?
- **Consistency**: Are rules internally consistent within each category?

### 2. Conflict Detection
- **Direct Conflicts**: Rules that directly contradict each other
- **Implicit Conflicts**: Rules that may conflict in certain scenarios
- **Priority Conflicts**: Cases where rule precedence is unclear

### 3. Duplicate Detection
- **Cursor System Prompt Duplicates**: Rules that duplicate capabilities already covered by Cursor's built-in system prompt
  - Common areas Cursor already handles: code quality, best practices, error handling, security awareness, performance optimization, clean code principles, proper variable naming, function structure, etc.
- **Inter-Rule Duplicates**: Rules within the user/project rules that have the same meaning but different wording
- **Redundant Requirements**: Multiple rules covering the same requirement or constraint
- **Overlapping Guidelines**: Rules with significant content overlap that could be consolidated

### 4. Gap Analysis (Optional)
Only mention if there are obvious missing areas that would significantly impact the project. It's perfectly fine if rules don't cover all areas - focus on what's actually needed:
- Code style and formatting
- Architecture and design patterns
- Testing requirements
- Documentation standards
- Security practices
- Performance considerations
- Error handling
- Code review process
- Deployment and CI/CD

### 5. Improvement Recommendations
For each identified issue, provide:
- **Problem**: Clear description of the issue
- **Impact**: Why this matters for the project
- **Solution**: Specific, actionable recommendation
- **Priority**: High/Medium/Low based on impact

## Output Format:
Please structure your response as follows:

### Executive Summary
[2-3 sentence overview of the rule set health]

### Critical Issues (High Priority)
[List of urgent issues]

### Conflicts Detected
[Detailed list of conflicting rules with recommendations]

### Duplicate Rules Found
[Rules that duplicate Cursor's built-in capabilities, with explanations of what Cursor already handles]
[Inter-rule duplicates with consolidation recommendations]

### Missing Rules (Gap Analysis - Optional)
[Only include if there are significant gaps that would impact the project]

### Improvement Opportunities
[Medium and low priority enhancements]

### Recommended Actions
[Prioritized list of next steps]

## Guidelines:
- Be constructive and solution-oriented
- Prioritize issues that could cause real development problems
- Consider the project context when making recommendations
- Suggest specific wording for new or modified rules
- Highlight rules that work well and should be maintained
- When identifying duplicates, provide consolidation suggestions that preserve the intent of all related rules
- **Cursor System Prompt Awareness**: Check if user rules duplicate what Cursor already does by default (e.g., code quality, best practices, error handling, security awareness. Refer to the Cursor system prompt that leaked in https://github.com/x1xhlol/system-prompts-and-models-of-ai-tools/blob/main/Cursor%20Prompts/Agent%20Prompt%20v1.0.txt)
- For Cursor system prompt duplicates, recommend removing redundant user rules and explain what Cursor already handles
- Consider semantic similarity, not just exact wording, when detecting duplicates`
}


